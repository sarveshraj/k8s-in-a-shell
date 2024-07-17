import { createServer } from 'http'
import fetch from 'fetch'

const contentTypeTextPlain = { 'Content-Type': 'text/plain' }

const pingHandler = (req, res) => {
  if (req.method !== 'GET') {
    res.writeHead(405, contentTypeTextPlain)
    res.end('Method Not Allowed')
    return
  }

  res.writeHead(200, contentTypeTextPlain)
  res.end('pong')
}

const homeHandler = (req, res) => {
  if (req.method !== 'GET') {
    res.writeHead(405, contentTypeTextPlain)
    res.end('Method Not Allowed')
    return
  }

  res.writeHead(200, { 'Content-Type': 'text/html' })
  res.end(`
      <form action="/submit-wage" method="post">
        <input type="text" name="employee_id" placeholder="Enter your employee id" required> 
        <input type="number" name="wage" placeholder="Enter your wage" required>
        <button type="submit">Pay tax</button>
      </form>
    `)
}

const submitWageHandler = (req, res) => {
  if (req.method !== 'POST') {
    res.writeHead(405, contentTypeTextPlain)
    res.end('Method Not Allowed')
    return
  }

  let body = ''
  req.on('data', chunk => {
    body += chunk.toString()
  })

  req.on('end', async () => {
    try {
      await fetch('server-service.learnk8s.svc.cluster.local:3000/paytax', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: body,
      })
      res.writeHead(200, contentTypeTextPlain)
      res.end('Wage saved successfully')
    } catch (error) {
      res.writeHead(500, contentTypeTextPlain)
      res.end(`Submission failed: ${error.message}`)
    }
  })
}

const server = createServer(async (req, res) => {
  switch (req.url) {
    case '/ping':
      pingHandler(req, res)
      break
    case '/':
      homeHandler(req, res)
      break
    case '/submit-wage':
      submitWageHandler(req, res)
      break
    default:
      res.writeHead(404, contentTypeTextPlain)
      res.end('Not Found')
  }
})

const port = 3000
server.listen(port, () => {
  console.log(`Server running at http://localhost:${port}/`)
})