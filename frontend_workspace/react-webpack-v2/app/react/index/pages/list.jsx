import React       from "react"
import $           from 'jquery'

import ListDetails from "../components/list/details.jsx"


export default class List extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
        list: []
    }

    this.getQuestionnaireList = this.getQuestionnaireList.bind(this)
    this.delBtnEvt = this.delBtnEvt.bind(this)
  }

  //Ajax
  callAjax(passedURL, method, id) {
    console.log("[List]:callAjax")

    let that = this
    //let method = 'get'
    let contentType = "application/json"

    $.ajax({
      url: encodeURI(passedURL),
      type: method,
      //cache    : false,
      crossDomain: false,
      contentType: contentType,
      dataType:    'json', //data type from server
      data:        ''
    }).done(function (data, textStatus, jqXHR) {
      console.log(data)
      if(method == 'get'){
        that.setState({
          list: data.list
        })
      }else if(method == 'delete'){
        //remove element
        let newList = that.state.list
        let delIndex = 0
        //search index
        for(let i=0,len=newList.length; i<len; i++){
            if(newList[i].id==id){
                delIndex = i
                break
            }
        }
        newList.splice(delIndex, 1)

        that.setState({
          list: newList
        })
      }
    }).fail(function (jqXHR, textStatus, errorThrown) {
      console.error(passedURL, textStatus, errorThrown.toString())
      swal("error!", "validation error was occurred!", "error")
    })
  }

  //Get Questionnaire List
  getQuestionnaireList() {
    console.log("[List]:getQuestionnaireList()")

    //call ajax
    //let url = '/admin/json/questionnaireList.json'
    let url = '/api/ques'
    this.callAjax(url, 'get', 0)
  }

  //Click delete btn
  delBtnEvt(id) {
    console.log("[List]:delBtnEvt()")

    let url = '/api/ques/'+id
    this.callAjax(url, 'delete', id)
  }

  //Only once before first render()
  componentWillMount() {
    console.log("[List]:componentWillMount()")
    this.getQuestionnaireList()
  }

  //
  render() {
    console.log("[List]render()")

    let that = this
    let list = this.state.list.map(function (data) {
      let key='questionnaire_' + data.id
      return (
        <ListDetails key={key} id={data.id} title={data.title} questions={data.questions} btnDel={that.delBtnEvt} />
      )
    })

    return (
      <div>
        <h1>List</h1>
        <table className="table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Title</th>
              <th>Answer</th>
            </tr>
          </thead>
          <tbody>
            {list}
          </tbody>
        </table>
      </div>
    )
  }
}
