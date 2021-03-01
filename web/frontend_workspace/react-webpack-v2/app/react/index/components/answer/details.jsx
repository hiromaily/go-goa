import React from "react"


export default class AnswerDetails extends React.Component {
  constructor(props) {
    super(props)

    this.changeAnswerEvt = this.changeAnswerEvt.bind(this)
 }

  //Change answer event
  changeAnswerEvt(e) {
    console.log("[AnswerDetails]:changeAnswerEvt(e)")
    //call
    this.props.chgA.call(this, e.target.value, this.props.index)
  }

  //
  render() {
    return (
      <tr>
        <th scope="row">{this.props.index+1}</th>
        <td>{this.props.question}</td>
        <td>
            <p className="text-danger">{this.props.error}</p>
            <div className="input-group">
              <input type="text" className="form-control" placeholder="answer"
               onChange={this.changeAnswerEvt} value={this.props.answer} aria-describedby="basic-addon2" />
            </div>
        </td>
      </tr>
    )
  }
}
