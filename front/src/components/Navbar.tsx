//import Avatar from "./Avatar";
import temporaryLogo from '../assets/temporarylogo.png'
import Avatar from '@mui/joy/Avatar';
import {Dropdown, ListDivider, ListItemDecorator, Menu, MenuItem} from "@mui/joy";
import MenuButton from "@mui/joy/MenuButton";
import React from "react";
import Settings from '@mui/icons-material/Settings';

const styles = {
    navbar: {
        display: 'flex',
        justifyContent: 'space-between',
        backgroundColor: '#fff',
        boxShadow: '0 1px 3px rgba(0, 0, 0, 0.1)',
        overflow: 'hidden',
        height: '50px',
        alignItems: 'center',
    },

    leftmost: {
        marginLeft: '35px',
    },

    rightmost: {
        cursor: 'pointer',
        marginRight: '35px',
    },

    menubutton: {
        border: 'none',
        padding: '0',
        borderRadius: '50%',
        cursor: 'pointer',
    }
}

export default function Navbar() {
    return (
        <div style={styles.navbar}>
            <div style={styles.leftmost}>
                logo
            </div>
            <div style={styles.rightmost}>
                <Dropdown>
                    <MenuButton style={styles.menubutton}>
                        <Avatar src={"/azavech.jpg"}>AD</Avatar>
                    </MenuButton>
                    <Menu
                        variant="solid"
                        invertedColors
                        aria-labelledby="apps-menu-demo"
                        placement="bottom-end"
                        sx={{'--ListItemDecorator-size': '24px'}}
                    >
                        <MenuItem orientation="horizontal">
                            <ListItemDecorator>
                                <Avatar>AD</Avatar>
                            </ListItemDecorator>
                            <div/>
                            Alice DI NUNNO
                        </MenuItem>
                        <MenuItem orientation="horizontal">
                            <ListItemDecorator>
                                <Avatar>TL</Avatar>
                            </ListItemDecorator>
                            <div/>
                            A
                        </MenuItem>
                        <MenuItem orientation="horizontal">
                            <ListItemDecorator>
                                <Avatar>NM</Avatar>
                            </ListItemDecorator>
                            <div/>
                            B
                        </MenuItem>
                        <MenuItem orientation="horizontal">
                            <ListItemDecorator>
                                <Avatar>DL</Avatar>
                            </ListItemDecorator>
                            <div/>
                            C
                        </MenuItem>
                        <MenuItem orientation="horizontal">
                            <ListItemDecorator>
                                <Avatar>💩</Avatar>
                            </ListItemDecorator>
                            <div/>
                            D
                        </MenuItem>
                        <ListDivider />
                        <MenuItem orientation="horizontal">
                            <ListItemDecorator>
                                <Avatar>+</Avatar>
                            </ListItemDecorator>
                            <div/>
                            E
                        </MenuItem>
                        <MenuItem orientation="horizontal">
                            <ListItemDecorator>
                                <Avatar><Settings/></Avatar>
                            </ListItemDecorator>
                            <div/>
                            F
                        </MenuItem>
                    </Menu>
                </Dropdown>
            </div>
        </div>
    )
}