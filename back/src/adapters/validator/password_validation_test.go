package validator

type testPasswordFieldValidator struct {
	password string `validate:"password"`
}

var validPassword = testPasswordFieldValidator{password: "Test123!"}
var validPasswordWithSpaces = testPasswordFieldValidator{password: "Test 123!"}
var invalidPasswordNoNumber = testPasswordFieldValidator{password: "Test_test"}
var invalidPasswordNoUppercase = testPasswordFieldValidator{password: "test123!"}
var invalidPasswordNoLowercase = testPasswordFieldValidator{password: "TEST123!"}
var invalidPasswordNoSpecialCharacter = testPasswordFieldValidator{password: "Test1234"}
var invalidPasswordTooShort = testPasswordFieldValidator{password: "Tt1!"}
var invalidPasswordTooLong = testPasswordFieldValidator{password: "azertyuiopqsdfghjklmwxcvbN123546789-azertyuiopqsdfghjklmwxcvbN123546789-"}
var invalidPasswordWithMultipleErrors = testPasswordFieldValidator{password: "test"}
var invalidEmptyPassword = testPasswordFieldValidator{password: ""}

func (suite *ValidationTestSuite) TestStringHasNumberAtStart() {
	// Given
	stringToTest := "7est"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasNumberInMiddle() {
	// Given
	stringToTest := "te5t"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasNumberAtEnd() {
	// Given
	stringToTest := "tes7"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasNoNumber() {
	// Given
	stringToTest := "test"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	suite.Assert().Equal(result, false)
}

func (suite *ValidationTestSuite) TestStringHasUppercaseLetterAtStart() {
	// Given
	stringToTest := "Test"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasUppercaseLetterInMiddle() {
	// Given
	stringToTest := "teSt"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasUppercaseLetterAtEnd() {
	// Given
	stringToTest := "tesT"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasNoUppercaseLetter() {
	// Given
	stringToTest := "test"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, false)
}

func (suite *ValidationTestSuite) TestStringHasLowercaseLetterAtStart() {
	// Given
	stringToTest := "tEST"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasLowercaseLetterInMiddle() {
	// Given
	stringToTest := "TEsT"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasLowercaseLetterAtEnd() {
	// Given
	stringToTest := "TESt"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasNoLowercaseLetter() {
	// Given
	stringToTest := "TEST"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	suite.Assert().Equal(result, false)
}

func (suite *ValidationTestSuite) TestStringHasSpecialCharacterAtStart() {
	// Given
	stringToTest := "!est"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasSpecialCharacterInMiddle() {
	// Given
	stringToTest := "te~t"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasSpecialCharacterAtEnd() {
	// Given
	stringToTest := "tes!"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	suite.Assert().Equal(result, true)
}

func (suite *ValidationTestSuite) TestStringHasNoSpecialCharacter() {
	// Given
	stringToTest := "test"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	suite.Assert().Equal(result, false)
}

func (suite *ValidationTestSuite) TestValidPassword() {
	//Given
	password := validPassword

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().True(ok)
	suite.Assert().Empty(errors)
}

func (suite *ValidationTestSuite) TestValidPasswordWithSpaces() {
	//Given
	password := validPasswordWithSpaces

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().True(ok)
	suite.Assert().Empty(errors)
}

func (suite *ValidationTestSuite) TestInvalidPasswordNoNumber() {
	//Given
	password := invalidPasswordNoNumber

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 1)
}

func (suite *ValidationTestSuite) TestInvalidPasswordNoUppercase() {
	//Given
	password := invalidPasswordNoUppercase

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 1)
}

func (suite *ValidationTestSuite) TestInvalidPasswordNoLowercase() {
	//Given
	password := invalidPasswordNoLowercase

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 1)
}

func (suite *ValidationTestSuite) TestInvalidPasswordNoSpecialCharacter() {
	//Given
	password := invalidPasswordNoSpecialCharacter

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 1)
}

func (suite *ValidationTestSuite) TestInvalidPasswordTooShort() {
	//Given
	password := invalidPasswordTooShort

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 1)
}

func (suite *ValidationTestSuite) TestInvalidPasswordTooLong() {
	//Given
	password := invalidPasswordTooLong

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 1)
}

func (suite *ValidationTestSuite) TestInvalidPasswordWithMultipleErrors() {
	//Given
	password := invalidPasswordWithMultipleErrors

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 4)
}

func (suite *ValidationTestSuite) TestInvalidEmptyPassword() {
	//Given
	password := invalidEmptyPassword

	//When
	ok, errors := suite.validator.Validate(password)

	//Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(errors)
	suite.Assert().Len(errors, 5)
}
