import { 
    BrowserRouter, 
    Routes, 
    Route 
} from "react-router-dom";
import {FrontPage} from './pages/FrontPage'
import {Monitor} from './pages/Monitor'

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
