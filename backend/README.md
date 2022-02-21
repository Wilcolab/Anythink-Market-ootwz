# Anythink Market Backend

The Anythink Market backend is Ruby web app written with [Ruby On Rails](https://rubyonrails.org/)

## Setup

Before running the app for the first time, make sure that:

- Ruby 2.7.0 is installed (follow the relevant instructions based on your Ruby version manager preference)
- Gems are installed (`bundle`)
- The DB is up and running
- The relvant DB exists. By default the DB is your username, so you can run `CREATE DATABASE <your username>;` in your favorite postgres client.
- Migrations have run (`bundle exec rails db:migrate`)

## Getting started

To start the app use: `./start.sh` from the backend directory.

## Dependencies

- [acts_as_follower](https://github.com/tcocca/acts_as_follower) - For implementing followers/following
- [acts_as_taggable](https://github.com/mbleigh/acts-as-taggable-on) - For implementing tagging functionality
- [Devise](https://github.com/plataformatec/devise) - For implementing authentication
- [Jbuilder](https://github.com/rails/jbuilder) - Default JSON rendering gem that ships with Rails, used for making reusable templates for JSON output.
- [JWT](https://github.com/jwt/ruby-jwt) - For generating and validating JWTs for authentication

## Folders

- `app/models` - Contains the database models for the application where we can define methods, validations, queries, and relations to other models.
- `app/views` - Contains templates for generating the JSON output for the API
- `app/controllers` - Contains the controllers where requests are routed to their actions, where we find and manipulate our models and return them for the views to render.
- `config` - Contains configuration files for our Rails application and for our database, along with an `initializers` folder for scripts that get run on boot.
- `db` - Contains the migrations needed to create our database schema.
