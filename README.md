# NowLive

Prototype of live theatre subscription management service. For theatre patrons
who want to manage season subscriptions or check the closing date of a show, a
responsive SMS-based messaging service could be appealing. The skeleton of how
one might implement one using the [Twilio](https://twilio.com) API is included
here, written in Golang.

This application is currently deployed on
[Heroku](https://theatre-now-live.herokuapp.com/), but
does not use a frontend, given that it is purely backend-based. In order to
set up the application for oneself, one would need to set the environment
variables `ACCOUNT_SID`, `AUTH_TOKEN` and `TWIL_PHONE` according to the
Twilio API instructions.

## Potential Improvements

- The service does not attempt any authentication of users, and a real-world deployment could leak private information
- Building an interactive method for using credits on tickets would be ideal
- Storing state in the database to improve interaction
- Setting up local deployments with Postgres and not Sqlite
- Building interface for easy data export / import for the theatre
