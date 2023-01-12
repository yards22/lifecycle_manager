import React from "react";
import ReactDOM from "react-dom/client";
import styled from "styled-components";
import App from "./App";
import ProvidedApp from "./ProvidedApp";

const SIndex = styled.div`
  display: flex;
  justify-content: center;
  height: calc(100vh - 60px);
  width: 100vw;
  margin-left: 10px;
  margin-right: 10px;
  margin-top: 65px;
`;

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <ProvidedApp>
      <SIndex>
        <App />
      </SIndex>
    </ProvidedApp>
  </React.StrictMode>
);
