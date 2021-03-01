import React from "react"
import { Link } from "react-router"


export default class ListDetail extends React.Component {
  constructor(props) {
    super(props)

    this.clickBtnEvt = this.clickBtnEvt.bind(this)
  }

  //Click del button
  clickBtnEvt(e) {
    console.log("[ListDetail]:clickBtnEvt(e)")

    //call event for post btn click
    this.props.btnDel.call(this, this.props.id)
  }

  //
  render() {

    return (
      <tr>
        <th scope="row">{this.props.id}</th>
        <td>{this.props.title}</td>
        <td>
          <Link to={"/answer/" + this.props.id} query={{t: this.props.title, q:this.props.questions}}>
            answer
          </Link>
        </td>
      </tr>
    )
  }
}
