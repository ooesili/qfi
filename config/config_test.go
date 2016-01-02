package config_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/ooesili/qfi/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		configDir      string
		badName        string
		errInvalidName string
	)

	BeforeEach(func() {
		var err error
		configDir, err = ioutil.TempDir("", "qfi-test-")
		Expect(err).ToNot(HaveOccurred())

		// simple symlink helper function
		symlink := func(dir, name string) {
			err := os.Symlink(dir, filepath.Join(configDir, name))
			Expect(err).ToNot(HaveOccurred())
		}

		symlink("/foo/bar", "foobar")
		symlink("/biz/baz", "bizbaz")
		symlink("/foo/bar/qux", "qux")

		// used often in input validation tests
		badName = fmt.Sprintf("bad%cname", os.PathSeparator)
		errInvalidName = fmt.Sprintf("invalid target name: %s", badName)
	})

	AfterEach(func() {
		err := os.RemoveAll(configDir)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("New", func() {
		Context("when the config directory containing non-symlink files", func() {
			It("returns an error", func() {
				badFile := filepath.Join(configDir, "notalink")
				file, err := os.Create(badFile)
				Expect(err).ToNot(HaveOccurred())
				file.Close()

				_, err = New(configDir)
				Expect(err).To(MatchError(
					fmt.Sprintf("not a symlink: %s", badFile),
				))
			})
		})

		Context("when a nonexistent directory is given", func() {
			It("creates the directory", func() {
				notadir := filepath.Join(configDir, "notadir")

				_, err := New(notadir)
				Expect(err).ToNot(HaveOccurred())
				Expect(notadir).To(BeADirectory())
			})
		})

		Context("when given a normal file is given", func() {
			It("returns an error", func() {
				badConfigDir := filepath.Join(configDir, "notadir")
				file, err := os.Create(badConfigDir)
				Expect(err).ToNot(HaveOccurred())
				file.Close()

				_, err = New(badConfigDir)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(HavePrefix(
					"cannot create directory: %s", badConfigDir))
			})
		})
	})

	Context("with a valid config struct", func() {
		var c *Config

		BeforeEach(func() {
			var err error
			c, err = New(configDir)
			Expect(err).ToNot(HaveOccurred())
		})

		Describe("Config.Resolve", func() {
			Context("when given existing targets", func() {
				It("can resolve their destinations", func() {
					destination, err := c.Resolve("foobar")
					Expect(err).ToNot(HaveOccurred())
					Expect(destination).To(Equal("/foo/bar"))

					destination, err = c.Resolve("bizbaz")
					Expect(err).ToNot(HaveOccurred())
					Expect(destination).To(Equal("/biz/baz"))

					destination, err = c.Resolve("qux")
					Expect(err).ToNot(HaveOccurred())
					Expect(destination).To(Equal("/foo/bar/qux"))

				})
			})

			Context("when given a nonexistent target", func() {
				It("returns an error", func() {
					_, err := c.Resolve("badtarget")
					Expect(err).To(MatchError("target 'badtarget' does not exist"))
				})
			})

			Context("when given an invalid target name", func() {
				It("returns an error", func() {
					_, err := c.Resolve(badName)
					Expect(err).To(MatchError(errInvalidName))
				})
			})
		})

		Describe("Config.Add", func() {
			It("can add a target when given an relative path", func() {
				// add target
				err := c.Add("newtarget", "/dingus")
				Expect(err).NotTo(HaveOccurred())

				// read the target's symlink
				destination, err := os.Readlink(filepath.Join(configDir, "newtarget"))
				Expect(err).NotTo(HaveOccurred())

				// make sure link points to the right place
				Expect(destination).To(Equal("/dingus"))
			})

			It("can add a target when given an relative path", func() {
				// add target
				err := c.Add("newtarget", "dingus")
				Expect(err).NotTo(HaveOccurred())

				// read the target's symlink
				destination, err := os.Readlink(filepath.Join(configDir, "newtarget"))
				Expect(err).NotTo(HaveOccurred())

				// make sure link points to the right place
				pwd, err := os.Getwd()
				Expect(err).NotTo(HaveOccurred())
				Expect(destination).To(Equal(filepath.Join(pwd, "dingus")))
			})

			Context("when config directory is read only", func() {
				BeforeEach(func() {
					err := os.Chmod(configDir, 0500)
					Expect(err).NotTo(HaveOccurred())
				})

				AfterEach(func() {
					err := os.Chmod(configDir, 0700)
					Expect(err).NotTo(HaveOccurred())
				})

				It("returns an error", func() {
					err := c.Add("newtarget", "dingus")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(HavePrefix("cannot create symlink: %s: ",
						filepath.Join(configDir, "newtarget")))
				})
			})

			Context("when given an invalid target name", func() {
				It("returns an error", func() {
					err := c.Add(badName, "/bleeblah")
					Expect(err).To(MatchError(errInvalidName))
				})
			})
		})

		Describe("Config.List", func() {
			It("lists all target names in alphabetical order", func() {
				Expect(c.List()).To(Equal([]string{
					"bizbaz", "foobar", "qux",
				}))
			})
		})

		Describe("Config.Delete", func() {
			Context("when given no targets", func() {
				It("does nothing", func() {
					err := c.Delete()
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("when given an existing target", func() {
				It("removes that target", func() {
					err := c.Delete("foobar")
					Expect(err).ToNot(HaveOccurred())

					_, err = os.Readlink(filepath.Join(configDir, "foobar"))
					Expect(err).To(HaveOccurred())
				})
			})

			Context("when given multiple existing targets", func() {
				It("removes all of the targets", func() {
					err := c.Delete("foobar", "bizbaz")
					Expect(err).ToNot(HaveOccurred())

					_, err = os.Readlink(filepath.Join(configDir, "foobar"))
					Expect(err).To(HaveOccurred())

					_, err = os.Readlink(filepath.Join(configDir, "bizbaz"))
					Expect(err).To(HaveOccurred())
				})
			})

			Context("when given a nonexistent target", func() {
				It("returns an error and does not delete any given targets", func() {
					err := c.Delete("foobar", "badtarget")
					Expect(err).To(MatchError("target 'badtarget' does not exist"))

					_, err = os.Readlink(filepath.Join(configDir, "foobar"))
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("when given an invalid target name", func() {
				It("returns an error", func() {
					err := c.Delete(badName)
					Expect(err).To(MatchError(errInvalidName))
				})
			})
		})

		Describe("Config.Move", func() {
			Context("when given an exsting target", func() {
				Context("when given an absolute path", func() {
					It("moves that target's destination", func() {
						err := c.Move("foobar", "/some/new/place")
						Expect(err).ToNot(HaveOccurred())

						destination, err := os.Readlink(filepath.Join(configDir, "foobar"))
						Expect(err).ToNot(HaveOccurred())
						Expect(destination).To(Equal("/some/new/place"))
					})
				})

				Context("when given an absolute path", func() {
					It("moves that target's destination", func() {
						err := c.Move("foobar", "relativetarget")
						Expect(err).ToNot(HaveOccurred())

						destination, err := os.Readlink(filepath.Join(configDir, "foobar"))
						Expect(err).ToNot(HaveOccurred())

						pwd, err := os.Getwd()
						Expect(err).ToNot(HaveOccurred())
						Expect(destination).To(Equal(filepath.Join(pwd, "relativetarget")))
					})
				})
			})

			Context("when given a nonexistent target", func() {
				It("returns an error", func() {
					err := c.Move("badtarget", "/what/evs")
					Expect(err).To(MatchError("target 'badtarget' does not exist"))
				})
			})

			Context("when given an invalid target name", func() {
				It("returns an error", func() {
					err := c.Rename("foobar", badName)
					Expect(err).To(MatchError(errInvalidName))
				})
			})
		})

		Describe("Config.Rename", func() {
			Context("when given an existing target", func() {
				It("renames the target", func() {
					err := c.Rename("foobar", "boofar")
					Expect(err).ToNot(HaveOccurred())

					_, err = os.Readlink(filepath.Join(configDir, "foobar"))
					Expect(err).To(HaveOccurred())

					destination, err := os.Readlink(filepath.Join(configDir, "boofar"))
					Expect(err).ToNot(HaveOccurred())
					Expect(destination).To(Equal("/foo/bar"))
				})

				Context("when the new name already exists", func() {
					It("returns an error", func() {
						err := c.Rename("foobar", "qux")
						Expect(err).To(MatchError("target 'qux' exists"))
					})
				})
			})

			Context("when given a nonexistent target", func() {
				It("returns an error", func() {
					err := c.Rename("badtarget", "boofar")
					Expect(err).To(MatchError("target 'badtarget' does not exist"))
				})
			})

			Context("when given an invalid target name for the old name", func() {
				It("returns an error", func() {
					err := c.Rename(badName, "foobar")
					Expect(err).To(MatchError(errInvalidName))
				})
			})

			Context("when given an invalid target name for the new name", func() {
				It("returns an error", func() {
					err := c.Rename("foobar", badName)
					Expect(err).To(MatchError(errInvalidName))
				})
			})
		})
	})
})
