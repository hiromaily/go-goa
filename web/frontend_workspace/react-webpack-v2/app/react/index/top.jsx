import React     from "react"
import ReactDOM  from "react-dom"
import { Router, Route, IndexRoute, hashHistory } from "react-router"

import Layout    from "./pages/layout.jsx"

import List      from "./pages/list.jsx"
import Answer    from "./pages/answer.jsx"


export default class App extends React.Component {
  render() {
    return (
      <Router history={hashHistory}>
        <Route path="/" component={Layout}>
          <IndexRoute component={List}></IndexRoute>
          <Route path="answer(/:id)" name="answer" component={Answer}></Route>
        </Route>
      </Router>
    )
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('root')
)
