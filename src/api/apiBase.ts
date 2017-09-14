const baseUrl = 'tasks-rad.quadient.com:8080';

function parseResponse(resolve, reject, response: Object) {
    if (response) {
        resolve(response);
    } else {
        reject();
    }
}

function callRequest<TRequest>(method: string, endPoint: string, params: Object): Promise<TRequest> {
    return new Promise((resolve, reject) => {
        const request = new XMLHttpRequest();
        request.addEventListener('load', () => parseResponse(resolve, reject, JSON.parse(request.response)));
        request.open(method, `http://${baseUrl}/${endPoint}`, true);
        request.withCredentials = true;
        request.send(JSON.stringify(params));
    });
}

export function callGet<TRequest>(path: string): Promise<TRequest> {
    return callRequest('GET', path, {});
}

export function callPut<TRequest>(path: string, params: Object): Promise<TRequest> {
    return callRequest('PUT', path, params);
}
