import React from "react"
import { Router } from "@reach/router"
import Admin from '../../components/admin/Admin'
let Home = () => <div>Dynamic App for cause owners goes here</div>

function App() {
  return (
    <Router>
      <Home path="app" />
      <Admin path="app/admin" />
    </Router>
  )
}

export default App
