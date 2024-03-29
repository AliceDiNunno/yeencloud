import React from 'react';
import {Stack, FormControl, FormLabel, Input, Button, DialogContent, Modal, DialogTitle, ModalDialog} from "@mui/joy";
 
interface FormPageProps {
    children: React.ReactNode
}

const FormPageComponent: React.FC<FormPageProps> = ({ children }) => {
    return (
        <div className="wrapper">
            <Modal open={true} ><div>
                <ModalDialog>
                    <DialogTitle>Create new project</DialogTitle>
                    <DialogContent>Fill in the information of the project.</DialogContent>
                    <form
                        onSubmit={(event: React.FormEvent<HTMLFormElement>) => {
                            event.preventDefault();
                        }}
                    >
                        <Stack spacing={2}>
                            <FormControl>
                                <FormLabel>Name</FormLabel>
                                <Input autoFocus required />
                            </FormControl>
                            <FormControl>
                                <FormLabel>Description</FormLabel>
                                <Input required />
                            </FormControl>
                            <Button type="submit">Submit</Button>
                        </Stack>
                    </form>
                </ModalDialog></div>
            </Modal>

            {children}
        </div>
    );
}

export default FormPageComponent;