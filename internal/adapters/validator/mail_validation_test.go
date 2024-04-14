package validator

type testMailFieldValidator struct {
	mail string `validate:"email"`
}

var standardMail = testMailFieldValidator{mail: "test@mailprovider.com"}
var validMailWithSubaddress = testMailFieldValidator{mail: "yeeencloud+test@mailprovider.com"}
var validMailWithPeriod = testMailFieldValidator{mail: "yeeencloud.test@mailprovider.com"}

var incorrectMail = testMailFieldValidator{mail: "testmail.com"}

func (suite *ValidationTestSuite) TestIncorrectEmail() {
	// Given
	mail := incorrectMail

	// When
	ok, err := suite.validator.Validate(mail)

	// Then
	suite.Assert().False(ok)
	suite.Assert().NotEmpty(err)
	suite.Assert().Len(err, 1)
}

func (suite *ValidationTestSuite) TestStandardEmail() {
	// Given
	mail := standardMail

	// When
	ok, err := suite.validator.Validate(mail)

	// Then
	suite.Assert().True(ok)
	suite.Assert().Empty(err)
}

func (suite *ValidationTestSuite) TestEmailWithSubaddress() {
	// Given
	mail := validMailWithSubaddress

	// When
	ok, err := suite.validator.Validate(mail)

	// Then
	suite.Assert().True(ok)
	suite.Assert().Empty(err)
}

func (suite *ValidationTestSuite) TestEmailWithPeriod() {
	// Given
	mail := validMailWithPeriod

	// When
	ok, err := suite.validator.Validate(mail)

	// Then
	suite.Assert().True(ok)
	suite.Assert().Empty(err)
}
