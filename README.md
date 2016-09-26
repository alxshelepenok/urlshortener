# URL Shortener

A simple library to shorten URL's with goo.gl written in Golang.

## Basic usage

	urlshortener.Setup(api_key)

	shortUrl, err := urlshortener.Shorten(longUrl)
	if err != nil {
		return "", err
	}

## Getting Api Key

**This is an excerpt from the official [documentation](https://developers.google.com/url-shortener/v1/getting_started).**

Requests to the Google URL Shortener API for public data must be accompanied by an identifier, which can be an [API key](https://developers.google.com/console/help/generating-dev-keys) or an [access token](https://developers.google.com/accounts/docs/OAuth2).

To acquire an API key:

+ Go to the [Google Developers Console](https://console.developers.google.com/).
+ Select a project, or create a new one.
+ In the sidebar on the left, expand **APIs & auth**. Next, click APIs. In the list of **APIs**, make sure the status is **ON** for the Google URL Shortener API.
+ In the sidebar on the left, select **Credentials**.
+ This API supports two types of credentials. Create whichever credentials are appropriate for your project:

    + **OAuth:** Your application must send an OAuth 2.0 token with any request that accesses private user data. Your application sends a client ID and, possibly, a client secret to obtain a token. You can generate OAuth 2.0 credentials for web applications, service accounts, or installed applications.
    </br>
    To create an OAuth 2.0 token, click **Create new Client ID**, provide the required information where requested, and click **Create Client ID**.

    + **Public API access:** A request that does not provide an OAuth 2.0 token must send an API key. The key identifies your project and provides API access, quota, and reports.
    </br>
    To create an API key, click **Create new Key** and select the appropriate key type. Enter the additional information required for that key type and click **Create**.

To keep your API keys secure, follow [the best practices for securely using API](https://developers.google.com/console/help/api-key-best-practices) keys.

After you have an API key, your application can append the query parameter `key=yourAPIKey` to all request URLs.

The API key is safe for embedding in URLs; it doesn't need any encoding.

## License
The MIT license.

Copyright (c) 2016 Alexandr Shelepenok [ashk.io](http://ashk.io)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
