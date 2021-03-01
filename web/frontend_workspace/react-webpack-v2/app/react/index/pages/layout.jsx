import React from "react"
import { Link } from "react-router"

import Footer from "../components/layout/footer.jsx"
import Nav from "../components/layout/nav.jsx"


export default class Layout extends React.Component {
  render() {
    console.log("[Layout]render()")

    const { location } = this.props
    const containerStyle = {
      marginTop: "60px"
    }

    return (
      <div>

        <Nav location={location} />

        <div className="container" style={containerStyle}>
          <div className="row">
            <div className="col-lg-12">
              <h1>Questionnaire For Users</h1>

              {this.props.children}

            </div>
          </div>
          <Footer/>
        </div>
      </div>

    )
  }
}
