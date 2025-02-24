# Exercises for chapter "RESTful API Development

## a) Simple API using net/http

Set up a simple web server using the standard-lib `net/http` package.

You **can** use a router such as https://github.com/julienschmidt/httprouter if you want.

Your web server should provide these features (combine them if possible):

1. understand a single query parameter e.g., http://localhost/products?keyword=pizza
2. understand multiple query parameters with the same name e.g., http://localhost/products?keyword=pizza&keyword=pasta
3. understand a path parameter e.g., http://localhost/products/P-123 ("P-123" is the path parameter)
4. provide different output depending on the provided `Accept` header
5. a simple POST request
6. a DELETE request

You can use some REST client like Postman or Insomnia for testing, or 
Intellij's built-in _HTTP Client_.

Things to skip / do-not-dos:

- error handling
- authentication / authorization
- other security aspects like CORS, CSRF
- API documentation
- endpoint testing

## b) Migration to a web framework

Now migrate your code to one of the popular web frameworks like Gin, Fibre or Echo.

You can work in small teams or pairs.

Take note of the challenges for later presentation to the group. 

What is your overall impression of your chosen framework, is it easy to use, powerful,
well documented?