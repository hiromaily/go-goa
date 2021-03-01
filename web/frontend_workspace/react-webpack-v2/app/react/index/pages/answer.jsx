import React     from "react"
import $         from 'jquery'

import AnswerDetails from "../components/answer/details.jsx"

export default class Answer extends React.Component {
  constructor(props) {
    super(props)
    //*
    let len = this.props.location.query.q.length
    let vals1 = new Array(len)
    let vals2 = new Array(len)
    //TODO:initialized variable have to prepare for respectively,
    //TODO:because "new array()"" has same addoress!!!!!
    for (let i=0; i<len; i++) {
      vals1[i] = ''
      vals2[i] = ''
    }
    this.state = {
        email: '',
        answers: vals1,
        errEmail: '',
        errAnswer: vals2,
    }
    //*/

    this.resetState = this.resetState.bind(this)
    this.changeEmailEvt = this.changeEmailEvt.bind(this)
    this.changeAnswer = this.changeAnswer.bind(this)
    this.createBtnEvt = this.createBtnEvt.bind(this)
  }

  //reset state
  resetState(){
    console.log("[Answer]:resetState")

    let len = this.props.location.query.q.length
    let vals1 = new Array(len)
    let vals2 = new Array(len)
    for (let i=0; i<len; i++) {
      vals1[i] = ''
      vals2[i] = ''
    }
    this.setState({
      email: '',
      answers: vals1,
      errEmail: '',
      errAnswer: vals2,
    })
  }

  //Ajax
  callAjax(passedURL, sendData) {
    console.log("[Answer]:callAjax")

    let that = this
    let method = 'post'
    let contentType = "application/json"
    if(sendData != ''){
      sendData = JSON.stringify(sendData);
    }

    $.ajax({
      url: encodeURI(passedURL),
      type: method,
      //cache    : false,
      crossDomain: false,
      contentType: contentType,
      dataType:    'json', //data type from server
      data:        sendData
    }).done(function (data, textStatus, jqXHR) {
      console.log(data)
      //TODO:reset form
      that.resetState()
      swal("success!", "thanks!", "success")
    }).fail(function (jqXHR, textStatus, errorThrown) {
      console.error(passedURL, textStatus, errorThrown.toString())
      swal("error!", "validation error was occurred or email has been already registered!", "error")
      //swal("error!", "validation error was occurred!", "error")
    })
  }

  //change email
  changeEmailEvt(e) {
    console.log("[Creation]:changeEmailEvt()")
    //console.log(e.target.value)

    this.setState({
      email: e.target.value,
    })
  }

  //change answer
  changeAnswer(value, index) {
    console.log("[Answer]:changeAnswer()")
    //console.log(value, index)

    let ans = this.state.answers
    ans[index] = value

    this.setState({
      answers: ans
    })
  }

  //check Email
  checkEmailAddress(str){
    if(str.match(/.+@.+\..+/)==null){
        return false;
    }else{
        return true;
    }
  }

  // validation
  checkValidation(){
    console.log("[Answer]:checkValidation()")
    //validation check
    let rtnVal = true

    let errEmail = ''
    if(this.state.email == ''){
      rtnVal = false
      errEmail = 'blank is not allowed'
    }else if(!this.checkEmailAddress(this.state.email)){
      rtnVal = false
      errEmail = 'email is invalid'
    }

    let errAnswer = this.state.answers.map(function (data, index) {
      if(data == ''){
        rtnVal = false
        return 'blank is not allowed'
      }else{
        return ''
      }
    })

    this.setState({
      errEmail : errEmail,
      errAnswer : errAnswer
    })
    return rtnVal
  }

  //Click create button
  createBtnEvt(e) {
    console.log("[Answer]:createBtnEvt(e)")

    //validation
    let bRet = this.checkValidation()
    console.log(bRet)
    if(!bRet){
      return
    }

    //call ajax [POST]
    let url = '/api/answer/'+this.props.params.id
    //"/api/answer/:id"

    //send data
    //{"email": "aaa@bbb.com", "answers":["a1","a2","a3"]}
    let sendData = new Object()
    sendData.email = this.state.email
    sendData.answers = this.state.answers

    console.log(sendData)
    this.callAjax(url, sendData)
  }

  //
  render() {
    console.log("[Answer]render()")
    //console.log(this.props.params.id)
    //console.log(this.props.location.query.q)

    console.log("state:", this.state)

    let that=this
    let questions = this.props.location.query.q.map(function (data, index) {
      let key='q_' + index
      let err = that.state.errAnswer[index]
      let answer = that.state.answers[index]
      //console.log("state is ", that.state)
      //console.log("err is ", err)
      return (
        <AnswerDetails key={key} index={index} question={data} answer={answer} chgA={that.changeAnswer} error={err} />
      )
    })

    return (
      <div>
        <h1>Answer</h1>
        <table className="table">
          <thead>
            <tr>
              <th>#</th>
              <th colSpan="2">{this.props.location.query.t}</th>
            </tr>
          </thead>
          <tbody>
            {questions}
          </tbody>
        </table>

        <p className="text-danger">{this.state.errEmail}</p>
        <div className="input-group">
          <span className="input-group-addon" id="sizing-addon2">E-Mail</span>
          <input type="text" className="form-control" placeholder="email@example.com"
          onChange={this.changeEmailEvt} aria-describedby="sizing-addon2"
          value={this.state.email} />
        </div>
        <br />
        <button className="btn btn-primary" onClick={this.createBtnEvt} type="button">Create</button>
        <br />

        <br />
      </div>
    )
  }
}
