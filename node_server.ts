const http = require("http");

const port = 8080;
console.log(`HTTP server listening on http://localhost:${port}`);

const server = http.createServer((req, res) => {
  let data = "";
  req.setEncoding("utf8");

  req.on("data", (chunk) => {
    data += chunk;
  });

  let json = {};
  req.on("end", () => {
    json = JSON.parse(data);

    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(JSON.stringify({
      data: json,
    }));
  });
});

server.listen(port);
