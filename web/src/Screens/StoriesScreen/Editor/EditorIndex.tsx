import React, { Component, useEffect, useState } from 'react';
import { EditorState, convertToRaw } from 'draft-js';
import SunEditor,{buttonList} from 'suneditor-react';
import 'suneditor/dist/css/suneditor.min.css';
import plugins from 'suneditor/src/plugins'


export default function EditorIndex(){
  const [editorState,setEditorState] = useState(EditorState.createEmpty())

  const onEditorStateChange= (editorState:any) => {
    setEditorState(editorState)
  };

  function handleChange(content:any){
    console.log("content",content); //Get Content Inside Editor
  }

  return (
       <SunEditor 
          height = "500px"
          placeholder="Please type here..."
          autoFocus={true}
          setOptions = {{
            height : "200px",
            plugins : plugins,
            buttonList: [
              [
                "formatBlock",
                "font",
                "fontSize",
                "fontColor",
                "align",
                "paragraphStyle",
                "blockquote"
              ],
              [
                "bold",
                "underline",
                "italic",
                "strike",
                "subscript",
                "superscript"
              ],
              ["removeFormat"],
              ["outdent", "indent"],
              ["table", "list"],
              ["link", "image", "video"]
            ]
          }}
          onChange = {handleChange}
       />
    );
}