const PutAuth = (url, postdata, cookie) => {
    fetch(url, {
            method: "PUT",
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": "Bearer " + cookie.token,
            }),
            body: JSON.stringify(postdata)
        }).then((data) => data.json())
}

export default PutAuth;