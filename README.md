# KOB - Knowledge Base Thingy

## Google OAUTH Setup

The application requires google oauth for logins.

1. Create a new application in https://console.developers.google.com
2. Using the credentials section create a new oauth client ID.
3. Set the applciation type to `Web Application`
4. For local development add authorized origin for `http://localhost:4200` and a redirect url for `http://localhost:4200/oauth/g/return`
5. Create a file in the project root called google-creds.json (configurable in the server flags).
6. Add the credentials to the file as `{"cid": "your client id", "csecret": "your client secret"}`


## Slack

1. Create a New App https://api.slack.com/apps
2. Enable Bot feature and install to your workspace.
3. Under OAuth & Permissions get the Bot User OAuth Access Token and use this as the `slack-token` cli flag.
4. Invite the bot to the channel.

