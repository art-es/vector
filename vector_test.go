package vector_test

import (
	"testing"

	"github.com/art-es/vector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vector Suite")
}

// TODO: add more complex cases

var _ = Describe("Vector", func() {
	Context(".Magnitude", func() {
		When("used simple two-dimensional vector with values 3 and 4", func() {
			v := vector.Vector{3, 4}
			It("result should be equal to 5", func() {
				Expect(v.Magnitude()).Should(Equal(5.0))
			})
		})
	})
})

var _ = Describe("Calculation functions", func() {
	Context("Add", func() {
		When("all requirements are met, addition of <2,2> and <3,3> vectors", func() {
			v1 := vector.Vector{2, 2}
			v2 := vector.Vector{3, 3}
			It("result should be equal to <5,5> without error", func() {
				res, err := vector.Add(v1, v2)
				Expect(err).Should(BeNil())
				Expect(res).Should(Equal(vector.Vector{5, 5}))
			})
		})

		When("used vectors with different sizes", func() {
			v1 := vector.Vector{0, 0}
			v2 := vector.Vector{0, 0, 0}
			It("result should be nil with error ErrDiffSizes", func() {
				res, err := vector.Add(v1, v2)
				Expect(err).Should(Equal(vector.ErrDiffSizes))
				Expect(res).Should(BeNil())
			})
		})
	})

	Context("Subtract", func() {
		When("all requirements are met, subtraction of <5,5> and <3,3> vectors", func() {
			v1 := vector.Vector{5, 5}
			v2 := vector.Vector{3, 3}
			It("result should be equal to <2,2> without error", func() {
				res, err := vector.Subtract(v1, v2)
				Expect(err).Should(BeNil())
				Expect(res).Should(Equal(vector.Vector{2, 2}))
			})
		})

		When("used vectors with different sizes", func() {
			v1 := vector.Vector{0, 0}
			v2 := vector.Vector{0, 0, 0}
			It("result should be nil with error ErrDiffSizes", func() {
				res, err := vector.Subtract(v1, v2)
				Expect(err).Should(Equal(vector.ErrDiffSizes))
				Expect(res).Should(BeNil())
			})
		})
	})

	Context("Multiply", func() {
		When("all requirements are met, multiplication of <2,2> and <3,3> vectors", func() {
			v1 := vector.Vector{2, 2}
			v2 := vector.Vector{3, 3}
			It("result should be equal to <6,6> without error", func() {
				res, err := vector.Multiply(v1, v2)
				Expect(err).Should(BeNil())
				Expect(res).Should(Equal(vector.Vector{6, 6}))
			})
		})

		When("used vectors with different sizes", func() {
			v1 := vector.Vector{0, 0}
			v2 := vector.Vector{0, 0, 0}
			It("result should be nil with error ErrDiffSizes", func() {
				res, err := vector.Multiply(v1, v2)
				Expect(err).Should(Equal(vector.ErrDiffSizes))
				Expect(res).Should(BeNil())
			})
		})
	})

	Context("Divide", func() {
		When("all requirements are met, dividation of <6,6> and <3,3> vectors", func() {
			v1 := vector.Vector{6, 6}
			v2 := vector.Vector{3, 3}
			It("result should be equal to <2,2> without error", func() {
				res, err := vector.Divide(v1, v2)
				Expect(err).Should(BeNil())
				Expect(res).Should(Equal(vector.Vector{2, 2}))
			})
		})

		When("used vectors with different sizes", func() {
			v1 := vector.Vector{0, 0}
			v2 := vector.Vector{0, 0, 0}
			It("result should be nil with error ErrDiffSizes", func() {
				res, err := vector.Divide(v1, v2)
				Expect(err).Should(Equal(vector.ErrDiffSizes))
				Expect(res).Should(BeNil())
			})
		})
	})
})

var _ = Describe("Dot & Cross Product functions", func() {
	Context("Dot", func() {
		When("all requirements are met, dot product of <1, 2> and <2, 3>", func() {
			v1 := vector.Vector{1, 2}
			v2 := vector.Vector{2, 3}
			It("result should be equal to 8 without error", func() {
				res, err := vector.DotProduct(v1, v2)
				Expect(err).Should(BeNil())
				Expect(res).Should(Equal(8.0))
			})
		})

		When("used vectors with different sizes", func() {
			v1 := vector.Vector{0, 0}
			v2 := vector.Vector{0, 0, 0}
			It("result should be zero with error ErrDiffSizes", func() {
				res, err := vector.DotProduct(v1, v2)
				Expect(err).Should(Equal(vector.ErrDiffSizes))
				Expect(res).Should(Equal(0.0))
			})
		})
	})

	Context("Cross", func() {
		When("all requirements are met, dot product of <1, 2, 3> and <2, 3, 4>", func() {
			v1 := vector.Vector{1, 2, 3}
			v2 := vector.Vector{2, 3, 4}
			It("result should be equal to <-1, 2, -1> without error", func() {
				res, err := vector.CrossProduct(v1, v2)
				Expect(err).Should(BeNil())
				Expect(res).Should(Equal(vector.Vector{-1, 2, -1}))
			})
		})

		When("used vectors which one of them is 4-dimensional vector", func() {
			v1 := vector.Vector{0, 0, 0}
			v2 := vector.Vector{0, 0, 0, 0}
			It("result should be nil with error ErrAvailableOnly3Dim", func() {
				res, err := vector.CrossProduct(v1, v2)
				Expect(err).Should(Equal(vector.ErrAvailableOnly3Dim))
				Expect(res).Should(BeNil())
			})
		})
	})
})
