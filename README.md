# sreeja_challenge
This repository contain the code and explanation for validating the credit card number.

# explanation
This code combines the functionalities from both parts:

isValidCreditCard function remains the same for validation.
handleRequest function:
Checks for POST requests.
Retrieves the credit card number from the form.
Validates the number using isValidCreditCard.
Writes the validation result to the response body.
main function:
Defines the handler for the root path (/).
Starts the web server on port 8080.
Important Note:

This is a very basic example and lacks functionalities for real-world deployment like:

Secure communication (HTTPS)
User interface (HTML form)
Integration with a database or external service
Remember, deploying a secure and scalable web application on AWS requires additional configuration and tools like VPC, EC2, and potentially CloudFormation.
