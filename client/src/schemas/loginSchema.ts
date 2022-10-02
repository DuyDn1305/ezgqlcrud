import { z } from "zod";

export const SignInSchema = z.object({
  email: z.string().min(1, "Field cannot empty"),
  password: z.string().min(1, "Field cannot empty")
});

export const SignUpSchema = z.object({
  name: z.string().min(1, "Must not empty"),
  email: z.string().email("Must be an email address"),
  password: z
    .string()
    .min(6, "Must contains at least 6 characters")
    .max(25, "Must have a maximum of 25 characters")
});
