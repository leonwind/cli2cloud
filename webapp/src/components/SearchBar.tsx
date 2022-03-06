import React, {ChangeEvent, Component, FormEvent, useState} from "react";
import styles from "../styles/SearchBar.module.css";
import {TextField} from "@mui/material"
import InputAdornment from '@material-ui/core/InputAdornment';
import SearchIcon from "@material-ui/icons/Search";
import { useNavigate } from "react-router-dom";

export const SearchBar = () => {
    const navigate = useNavigate();
    const [clientID, setClientID] = useState("");

    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        navigate("/" + clientID);
        window.location.reload();
    }

    return (
        <form onSubmit={handleSubmit}>
            <TextField variant="outlined"
                value={clientID} 
                onChange={e => setClientID(e.target.value)} type="text" 
                placeholder="Enter your client ID..."
                InputProps={{
                    endAdornment: (
                        <InputAdornment position="start">
                            <SearchIcon/>
                        </InputAdornment>
                    )}}
            /> 
        </form>
    )
}