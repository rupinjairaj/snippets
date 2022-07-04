export const endpoints = {
    "local": "http://localhost:8000",
    "dev": "https://snippet-web-dev.herokuapp.com",
    "prod": "https://snippet-web-prod.herokuapp.com"
}
export const snippetPath = "/snippet"
export const tagsPath = "/tag"

export const getURL = (hostname) => {
    if (hostname.includes("local")) {
        return endpoints['local']
    }
    if (hostname.includes("prod")) {
        return endpoints['prod']
    }
    if (hostname.includes("dev")) {
        return endpoints['dev']
    }
}
