import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '15s', target: 500 },       // duração em que o número de usuários irá atingir o limite alvo
    // { duration: '1m10s', target: 1000 },  // pode ser adicionar vários intervalos
    // { duration: '30s', target: 1000 },
  ],
  // vus: 20,                     // usuários simultâneos
  // duration: '50s',             // duração da execução
  insecureSkipTLSVerify: true,    // ignorar certificado TLS
};


export default function () {
  const res = http.get('http://localhost:3001/hello');
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
}
