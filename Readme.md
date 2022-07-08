# Simple uptime montioring

Simple uptime monitoring is a small tool to monitor the availability of websites, written in Go.

## In a nutshell

You can monitor the availability of any website with this tool. Just add some details about your website in the `config.yml` file and then execute the program. The program will the reach out to all endpoints you configured in `config.yml` and notifies you via email and/or Slack if the request results in an error.

## Configuration

The program of course needs to know which websites and routes it should monitor. This is where the `config.yml` comes into play. There, you can define the following:

- The website(s) you want to monitor (e.g., "mywebiste.com")
- Which specific urls of those websites should be pinged (e.g "/login". "/imprint" etc.)
- The request method for each url (POST, GET etc.)
- The headers that will be sent on requests (e.g., Authorization headers)

## Environment variables

Execute the following command from the command line:

```
cp .env.example
```

The values defined in the `.env` files will be important if you want to be notified about outages via email or Slack.

### Config.yml structure

The configuration is written in YAML format. The websites that will be monited are called "destinations" and thus need to be declared under the `destinations` property in `config.yml`.
The destinations property expects an array of objects, where each of those objects represents a website you want to monitor.

### Adding a new website to config.yml

Before you start reading this, open the `config.yml` file and try to understand its structure. It should already tell you almost everything you need to know in order to add your own configuration.

When you add a new website, the following this are required:

- `baseUrl`: The base url of the website you want to ping. For example, "https://google.com"
- `routes`: Routes expect an array of object that specify an route of the base url you want to ping (like slug, request method and request headers.). We will look at this in more detail in the next section

#### Routes

With routes you have granular control about the urls you want to monitor. Here's how to configure routes:

- `path`: The path of the route. For example, "/imprint". When running the program, the path will appended to the `baseUrl` you configured. So, if your `baseUrl` is "https://google.com", our example from before will lead to the URL "https://google.com/imprint".
- `method`: The HTTP request method in all capital letters. For example, `GET` will lead to `baseUrl/path` requested using the HTTP GET method.
- `headers`: An array of strings. Each sttring will be added as a header to the request you want to send. Use a colon (`:`) to separate key and value of the header you want to add. For example, `Authorization:Bearer 123346`. This property is especially useful if you want to monitor resources that require authorization.

## Notifications

Monitoring is pretty useless if you are not informed about errors. Thus, the program includes two notification methods: Email and Slack.

### Email

The SMTP settings need to be configured in the `.env` file. You can refer to the `.env.example` file to find the variables that need to be configured (the best way is to execute `cp .env.example .env` to make sure everything is preconfigured).

To ebable email support, make `EMAILS_ENABLED=true` in the `.env`.

### Slack

You can also be notified via Slack if you want. To enable this, populate the `SLACK_WEBHOOK_URL` in `.env` and set `SLACK_ENABLED=true`. For a tutorial on how to create Slack Webhooks go [here](https://api.slack.com/messaging/webhooks).

## To-Dos

- Make status code configurable: At the moment, an error will be sent whenever the response status code is not 200. It would be better, however, if this was only the default behaviour and would be configurable in the config file.
- Validate request method. At the moment, if is not checked if the `method` propery in `routes` is actually a valid HTTP method.
