import { 
    BrowserRouter, 
    Routes, 
    Route 
} from "react-router-dom";
import {FrontPage} from './components/FrontPage'
import {Monitor} from './components/OutputMonitor'

export function App (){
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<FrontPage />} />
                <Route path="/:clientID" element={<Monitor />} />
            </Routes> 
        </BrowserRouter>         
    );
}