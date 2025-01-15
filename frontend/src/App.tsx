
// import React from "react";

// import { BrowserRouter as Router } from "react-router-dom";

// import ConfigRoutes from "./routes";

// import "./App.css";

// import Payment  from '../src/pages/payment/payment'


// const App: React.FC = () => {

//   return (

//     <Router>
    
//       <Router path="/" element={<Payment />} />

//     </Router>

//   );

// };


// export default App;



import React from "react";
import { BrowserRouter as Router } from "react-router-dom";
import ConfigRoutes from "./routes";
import "./App.css";

const App: React.FC = () => {
  return (
    <Router>
      <ConfigRoutes />
    </Router>
  );
};

export default App;
