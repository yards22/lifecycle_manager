import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import ProvidedApp from "./ProvidedApp";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <ProvidedApp>
      <App />
    </ProvidedApp>
  </React.StrictMode>
);
