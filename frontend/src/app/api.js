
let Path = '';

export function setPath(path) {
    Path = path;
}

/* 
    Темы: "dark", "light"
    Размеры: "small", "normal", "big"
*/
export function initTheme() {
    const curTheme = localStorage.getItem("theme") || "light";
    document.documentElement.setAttribute("data-theme", curTheme);

    const curSize = localStorage.getItem("size") || "normal";
    document.documentElement.setAttribute("data-size", curSize);
}

export function setTheme(theme) {
    document.documentElement.setAttribute("data-theme", theme);
    localStorage.setItem("theme", theme);
}

export function setSize(size) {
    document.documentElement.setAttribute("data-size", size);
    localStorage.setItem("size", size);
}

export function module(cmd, args, callback, errfunc) {
    axios
        .post(Path + '/' + cmd, args)
        .then((response) => {
            if (response.data.error) {
                errfunc(response.data.error)
            } else {
                let answer = callback(response.data)
                if (answer) {
                    errfunc(answer)
                }
            }
        })
        .catch((error) => {
            errfunc(error)
        });
}

export function delapi(cmd, args, callback, errfunc) {
    axios
        .delete('api/' + cmd, args)
        .then((response) => {
            if (response.data.error) {
                errfunc(response.data.error)
            } else {
                let answer = callback(response.data)
                if (answer) {
                    errfunc(answer)
                }
            }
        })
        .catch((error) => {
            errfunc(error)
        });
}

export function postapi(cmd, args, callback, errfunc) {
    axios
        .post('api/' + cmd, args)
        .then((response) => {
            if (response.data.error) {
                errfunc(response.data.error)
            } else {
                let answer = callback(response.data)
                if (answer) {
                    errfunc(answer)
                }
            }
        })
        .catch((error) => {
            errfunc(error)
        });
}

export function putapi(cmd, args, callback, errfunc) {
    axios
        .put('api/' + cmd, args)
        .then((response) => {
            if (response.data.error) {
                errfunc(response.data.error)
            } else {
                let answer = callback(response.data)
                if (answer) {
                    errfunc(answer)
                }
            }
        })
        .catch((error) => {
            errfunc(error)
        });
}

export function api(cmd, args, callback, errfunc) {
    module('api/' + cmd, args, callback, errfunc)
}


export function getapi(cmd, callback, errfunc) {
    axios
        .get('api/' + cmd)
        .then((response) => {
            if (response.data.error) {
                errfunc(response.data.error)
            } else {
                let answer = callback(response.data)
                if (answer) {
                    errfunc(answer)
                }
            }
        })
        .catch((error) => {
            errfunc(error)
        });
}

// Dispatch event on click outside of node 
export function clickOutside(node) {

    const handleClick = event => {
        if (node && !node.parentElement.contains(event.target) && !event.defaultPrevented) {
            node.dispatchEvent(
                new CustomEvent('clickoutside', node)
            )
        }
    }

    document.addEventListener('click', handleClick, true);

    return {
        destroy() {
            document.removeEventListener('click', handleClick, true);
        }
    }
}

export function form(cmd, files, args, callback, errfunc, api) {
    let formData = new FormData();

    if (api) {
        cmd = 'api/' + cmd;
        for (const [key, value] of Object.entries(files)) {
            for (let i = 0; i < value.length; i++) {
                formData.append(key, value[i]);
            }
        }
    } else {
        for (let i = 0; i < files.length; i++) {
            formData.append("files", files[i]);
        }
    }
    formData.append("json", JSON.stringify(args));

    axios
        .post(Path + "/" + cmd, formData, {
            headers: {
                "Content-Type": "multipart/form-data",
            },
        })
        .then((response) => {
            if (response.data.error) {
                errfunc(response.data.error)
            } else {
                let answer = callback(response.data)
                if (answer) {
                    errfunc(answer)
                }
            }
        })
        .catch((error) => {
            errfunc(error)
        });
}

export function moduleform(cmd, files, args, callback, errfunc) {
    let formData = new FormData();

    for (const [key, value] of Object.entries(files)) {
        for (let i = 0; i < value.length; i++) {
            formData.append(key, value[i]);
        }
    }
    formData.append("json", JSON.stringify(args));

    axios
        .post(Path + "/" + cmd, formData, {
            headers: {
                "Content-Type": "multipart/form-data",
            },
        })
        .then((response) => {
            if (response.data.error) {
                errfunc(response.data.error)
            } else {
                let answer = callback(response.data)
                if (answer) {
                    errfunc(answer)
                }
            }
        })
        .catch((error) => {
            errfunc(error)
        });
}
