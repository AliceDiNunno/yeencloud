package validator

type testStringFieldValidator struct {
	field string `validate:"required"`
}

var emptyString = testStringFieldValidator{field: ""}
var nonEmptyString = testStringFieldValidator{field: "test"}

func (suite *ValidationTestSuite) TestRequiredStringIsEmpty() {
	// Given
	stringToTest := emptyString

	// When
	ok, errors := suite.validator.Validate(stringToTest)

	// Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 1)
}

func (suite *ValidationTestSuite) TestRequiredStringIsNonEmpty() {
	// Given
	stringToTest := nonEmptyString

	// When
	ok, errors := suite.validator.Validate(stringToTest)

	// Then
	suite.Assert().True(ok)
	suite.Assert().Empty(errors)
}
