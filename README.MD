# Saakaru Quest Bot

Saakaru Quest Bot a tool for auto claim Saakaru daily quest.


## Prerequisite
- go installed (im using version 1.22.2)


## Installation

1. Clone the repository:

```bash
git clone https://github.com/Widiskel/saakuru-bot.git
```

2. Install dependencies:

```bash
cd saakaru-bot
npm install
```

## Configuration

1. open `utils/accounts.go` file
2. Add your Saakaru Bearer tokens to the `accounts.go` file in the following format:

```go
package accounts

var Accounts = []string{
	"YOUR TOKEN HERE",
}

```

### Accessing Authorization Token

To access the Authorization token required for the `accounts.go` file, follow these steps:

1. Open your web browser and go to the Saakaru website.
2. Log in / Register to [Sakaru](https://gainz.saakuru.com/dashboard) by connecting your twitter account.
3. Enter invite code (You can use my invite code `X1OFZRL9`) or choose Skip & Continue.
4. Open the developer tools by pressing `F12` or right-clicking on the page and selecting "Inspect" or "Inspect Element".
5. Navigate to the "Network" tab.
6. Perform a request that requires authentication, such as refreshing the page or accessing a protected resource.
7. Look for the request in the list of network requests and find `refresh-token` request.
8. Click on the request to view its details.
9. In the `response` section, you should see something like this

```json
{
    "code": 0,
    "data": {
        "token": "YOURTOKEN",
        "refreshToken": "YOURREFRESHTOKEN"
    },
    "timestamp": "2024-04-27T11:55:33.759050876Z"
}
```

Copy that token value and paste it into the `accounts.go` file.

## Usage

To run the bot, execute the following command:

```bash
go run main.go
```

## Note

Some Quest like follow someone, connect wallet, or add address quest need to be completed manually
