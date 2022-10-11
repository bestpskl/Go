import http from 'k6/http';

// set option run k6
export let options = {
    vus: 5,
    duration: '5s'
}

export default function () {
    http.get('http://host.docker.internal:8000/products');

}