import { 
    BrowserRouter, 
    Routes, 
    Route 
} from "react-router-dom";
import {FrontPage} from './components/FrontPage'
import {OutputMonitor} from './components/OutputMonitor'

export function App (){
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<FrontPage />} />
                <Route path="/:clientID" element={<OutputMonitor />} />
            </Routes> 
        </BrowserRouter>         
    );
}