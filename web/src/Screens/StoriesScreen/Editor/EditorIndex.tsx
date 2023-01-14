// import RichTextEditor from '@mantine/rte';
// import React, { Component, useEffect, useState } from 'react';


// export default function EditorIndex(){
//   const [storyContent,setStoryContent] = useState('<p>Rich text editor content</p>') 

//   function handleChange(content:any){
//     console.log("content",content); //Get Content Inside Editor
//   }

//   return (
//      <RichTextEditor value={storyContent} onChange={setStoryContent} id="rte"/>
//     );
// }

import React, { Component } from 'react';
import { Editor } from 'react-draft-wysiwyg';
import '../node_modules/react-draft-wysiwyg/dist/react-draft-wysiwyg.css';


export const EditorIndex = () => <Editor />
