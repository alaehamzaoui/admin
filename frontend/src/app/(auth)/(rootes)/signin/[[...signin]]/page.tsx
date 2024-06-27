"use client"
import React, { useState } from 'react'
import * as z from "zod"
import { Button } from "@/components/ui/button"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from 'react-hook-form'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import Link from 'next/link'
import { BACKEND_URL } from '@/loadEnv';
import router from 'next/router'
import Modal from 'react-modal';

// Stile für das Modal
const customStyles = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-50%',
    transform: 'translate(-50%, -50%)',
    backgroundColor: '#f0f4f8',
    padding: '20px',
    borderRadius: '10px',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
  },
  overlay: {
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
  },
};

const signInSchema = z.object({
    email: z.string().email("Email must be valid."),
    password: z.string().min(6, "Password Should have at least 6 characters."),
})

const Page = () => {
    const [modalIsOpen, setModalIsOpen] = useState(false);
    const [modalMessage, setModalMessage] = useState('');

    const form = useForm<z.infer<typeof signInSchema>>({
        resolver: zodResolver(signInSchema),
        defaultValues: {
            email: "",
            password: "",
        },
    })

    async function onSubmit(data: z.infer<typeof signInSchema>) {
        try {
            const loginResponse = await fetch(BACKEND_URL + "/auth/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    email: data.email,
                    password: data.password,
                }),
            });
            if (loginResponse.ok) {
                const loginData = await loginResponse.json();
                localStorage.setItem("authToken", loginData.token);
                window.location.href = "/home";
            } else {
                const errorData = await loginResponse.json();
                if ( errorData.message == "ERR_INVALID_CREDENTIALS")
                setModalMessage("Email Oder Password ist falsch");
                else
                setModalMessage("Failed to login"); 
                setModalIsOpen(true);
            }
        } catch (error) {
            console.error('Failed to submit data', error);
            setModalMessage('Failed to submit data');
            setModalIsOpen(true);
        }
    }

    return (
        <>
            <div className="signUpWrapper">
                <div className="formWrapper">
                    <div className="left">
                        <h3 className="title">Hallo, Freunde!</h3>
                        <p>Geben Sie Ihre persönlichen Daten ein und starten Sie Ihre Reise mit uns</p>
                        <Link href={"/signup"}>
                            <Button className='border-zinc-500 text-white-300 hover:border-zinc-200 hover:text-white-100 hover:bg-teal-600 transition-colors border rounded-full px-8 bg-teal-600'>Sign Up</Button>
                        </Link>
                    </div>
                    <div className="right">
                        <h3 className='text-center text-2xl font-semibold'>Hier anmelden</h3>
                        <Form {...form}>
                            <form onSubmit={form.handleSubmit(onSubmit)}>
                                <FormField
                                    control={form.control}
                                    name="email"
                                    render={({ field }) => (
                                        <FormItem className='space-y-0 mb-2'>
                                            <FormLabel>Email</FormLabel>
                                            <FormControl>
                                                <Input placeholder="admin@example.com" {...field} />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                                <FormField
                                    control={form.control}
                                    name="password"
                                    render={({ field }) => (
                                        <FormItem className='space-y-0 mb-2'>
                                            <FormLabel>Password</FormLabel>
                                            <FormControl>
                                                <Input placeholder="********" type='password' {...field} />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                                <Button type="submit" className='w-full bg-teal-600 text-zinc-200 hover:border-zinc-200 hover:text-zinc-100 hover:bg-teal-600'>Submit</Button>
                            </form>
                        </Form>
                    </div>
                </div>
            </div>
            <Modal
                shouldCloseOnOverlayClick={false}
                isOpen={modalIsOpen}
                onRequestClose={() => setModalIsOpen(false)}
                style={customStyles}
                contentLabel="Error Message"
            >
                <h2 style={{ color: '#e74c3c', marginBottom: '10px' }}>Error</h2>
                <div style={{ marginBottom: '20px', fontSize: '16px', color: '#2c3e50' }}>{modalMessage}</div>
                <Button onClick={() => setModalIsOpen(false)} className='bg-teal-600 text-zinc-200'>Close</Button>
            </Modal>
        </>
    )
}

export default Page
