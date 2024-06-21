import express from 'express'
import cors from 'cors'
import helmet from 'helmet'
import morgan from 'morgan'
import { renderFile } from 'ejs'

const name = process.env.NAME || 'Jumaima'
const environment = process.env.ENVIRONMENT || 'development'

export const app = express()

app.use(express.static('./assets'))
app.set('view engine', 'ejs')
app.engine('html', renderFile)
app.set('views', 'views')

app.use(cors())
app.use(helmet())
app.use(express.json())
app.use(morgan(':method :url :status :res[content-length] - :response-time ms'))

app.get('/', (req, res) => {
  // Render the HTML file and pass the environment variable
  res.render('index.html', { username: name, environment: environment })
})
