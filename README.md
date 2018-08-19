# Fetch Distance Function

[![Build Status](https://travis-ci.org/bense4ger/unchained.fetchdistance.svg?branch=master)](https://travis-ci.org/bense4ger/unchained.fetchdistance)

This function fetches the distance between two points from the Goole Distance Matrix API

## Environment Variables and Flags

### Environment Varaibles

#### Env

`dev` | `prod`
Defaults to `prod`.  

#### API_KEY

The Google API key

### Command Line Flags

#### origin

`--origin=<origin>`
The origin location.  Should be a lat lon.

#### dest

`--dest=<destination>`
The destination location.  Should be a lat lon.

## Running Locally

The function can be run locally for testing/debugging/development.

`ENV=dev API_KEY=<api-key> [go run main.go | ./unchained.fetchdistance] --origin=<latlon> --dest=<latlon>`  

Note - running locally bypasses the Lambda handler function

## Deployment

The function is deployed on every commit to master.