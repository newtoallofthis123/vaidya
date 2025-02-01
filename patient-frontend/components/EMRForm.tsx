"use client";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Checkbox } from "@/components/ui/checkbox";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Label } from "@/components/ui/label";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import { AutoExpandingInputGroup } from "./AutoExpandingInputGroup";

const patientSchema = z.object({
  patient_id: z.string().min(1, "Patient ID is required"),
  first_name: z.string().min(1, "First name is required"),
  last_name: z.string().min(1, "Last name is required"),
  age: z.number().min(0).max(150),
  gender: z.enum(["male", "female", "other"]),
  address: z.string().min(1, "Address is required"),
  identity: z.string().min(1, "Identity is required"),
  phone: z.string().regex(/^\+?[1-9]\d{1,14}$/, "Invalid phone number"),
  email: z.string().email("Invalid email address"),
  description: z.string(),
  recurring: z.boolean(),
  problems: z.array(z.string()),
  medicines: z.array(z.string()),
  conditions: z.array(z.string()),
  diagnosis: z.string(),
  next_session: z
    .string()
    .regex(/^\d{4}-\d{2}-\d{2}$/, "Invalid date format (YYYY-MM-DD)"),
});

export type PatientFormData = z.infer<typeof patientSchema>;

interface EMRFormProps {
  initialData?: Partial<PatientFormData>;
  onSubmit: (data: PatientFormData) => void;
}

export function EMRForm({ initialData, onSubmit }: EMRFormProps) {
  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<PatientFormData>({
    resolver: zodResolver(patientSchema),
    defaultValues: {
      ...initialData,
      problems: initialData?.problems || [""],
      medicines: initialData?.medicines || [""],
      conditions: initialData?.conditions || [""],
    },
  });

  return (
    <Card className="w-full max-w-2xl mx-auto">
      <CardHeader>
        <CardTitle>Electronic Medical Record</CardTitle>
        <CardDescription>Enter patient information</CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div>
              <Label htmlFor="patient_id">Patient ID</Label>
              <Input id="patient_id" {...register("patient_id")} />
              {errors.patient_id && (
                <p className="text-red-500 text-sm">
                  {errors.patient_id.message}
                </p>
              )}
            </div>
            <div>
              <Label htmlFor="first_name">First Name</Label>
              <Input id="first_name" {...register("first_name")} />
              {errors.first_name && (
                <p className="text-red-500 text-sm">
                  {errors.first_name.message}
                </p>
              )}
            </div>
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div>
              <Label htmlFor="last_name">Last Name</Label>
              <Input id="last_name" {...register("last_name")} />
              {errors.last_name && (
                <p className="text-red-500 text-sm">
                  {errors.last_name.message}
                </p>
              )}
            </div>
            <div>
              <Label htmlFor="age">Age</Label>
              <Input
                id="age"
                type="number"
                {...register("age", { valueAsNumber: true })}
              />
              {errors.age && (
                <p className="text-red-500 text-sm">{errors.age.message}</p>
              )}
            </div>
          </div>
          <div>
            <Label>Gender</Label>
            <Controller
              name="gender"
              control={control}
              render={({ field }) => (
                <RadioGroup
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                  className="flex space-x-4"
                >
                  <div className="flex items-center space-x-2">
                    <RadioGroupItem value="male" id="male" />
                    <Label htmlFor="male">Male</Label>
                  </div>
                  <div className="flex items-center space-x-2">
                    <RadioGroupItem value="female" id="female" />
                    <Label htmlFor="female">Female</Label>
                  </div>
                  <div className="flex items-center space-x-2">
                    <RadioGroupItem value="other" id="other" />
                    <Label htmlFor="other">Other</Label>
                  </div>
                </RadioGroup>
              )}
            />
            {errors.gender && (
              <p className="text-red-500 text-sm">{errors.gender.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="address">Address</Label>
            <Textarea id="address" {...register("address")} />
            {errors.address && (
              <p className="text-red-500 text-sm">{errors.address.message}</p>
            )}
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div>
              <Label htmlFor="identity">Identity</Label>
              <Input id="identity" {...register("identity")} />
              {errors.identity && (
                <p className="text-red-500 text-sm">
                  {errors.identity.message}
                </p>
              )}
            </div>
            <div>
              <Label htmlFor="phone">Phone</Label>
              <Input id="phone" type="tel" {...register("phone")} />
              {errors.phone && (
                <p className="text-red-500 text-sm">{errors.phone.message}</p>
              )}
            </div>
          </div>
          <div>
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" {...register("email")} />
            {errors.email && (
              <p className="text-red-500 text-sm">{errors.email.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="description">Description</Label>
            <Textarea id="description" {...register("description")} />
            {errors.description && (
              <p className="text-red-500 text-sm">
                {errors.description.message}
              </p>
            )}
          </div>
          <div className="flex items-center space-x-2">
            <Checkbox id="recurring" {...register("recurring")} />
            <Label htmlFor="recurring">Recurring Patient</Label>
          </div>
          <AutoExpandingInputGroup
            control={control}
            name="problems"
            label="Problems"
            error={errors.problems}
          />
          <AutoExpandingInputGroup
            control={control}
            name="medicines"
            label="Medicines"
            error={errors.medicines}
          />
          <AutoExpandingInputGroup
            control={control}
            name="conditions"
            label="Conditions"
            error={errors.conditions}
          />
          <div>
            <Label htmlFor="diagnosis">Diagnosis</Label>
            <Textarea id="diagnosis" {...register("diagnosis")} />
            {errors.diagnosis && (
              <p className="text-red-500 text-sm">{errors.diagnosis.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="next_session">Next Session</Label>
            <Input
              id="next_session"
              type="date"
              {...register("next_session")}
            />
            {errors.next_session && (
              <p className="text-red-500 text-sm">
                {errors.next_session.message}
              </p>
            )}
          </div>
          <Button type="submit" className="w-full">
            Submit
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
