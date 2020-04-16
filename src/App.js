import React from 'react';
import Header from './components/Header';
import Main from './components/Main';
import { ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import './App.css';

function App() {
  return (
    <div>
      <Header />
      <Main />
      <ToastContainer
              position="top-right"
              autoClose={5000}
              hideProgressBar={false}
              newestOnTop={false}
              closeOnClick={true}
              pauseOnVisibilityChange={true}
              draggable={false}
              pauseOnHover={true}
            />
    </div>
  );
}

export default App;
