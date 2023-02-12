const DeleteAuth = (url, cookie) => {
    fetch(url, {
            method: "DELETE",
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": "Bearer " + cookie.token,
            })
        }).then((data) => data.json())
}

export default DeleteAuth;