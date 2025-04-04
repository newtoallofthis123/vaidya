"use client";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import * as z from "zod";
import { AutoExpandingInputGroup } from "./AutoExpandingInputGroup";

const patientSchema = z.object({
  patient_id: z.string().min(1, "Patient ID is required"),
  name: z.string().min(1, "First name is required"),
  age: z.number().min(0).max(150),
  gender: z.enum(["Male", "Female", "Other"]),
  address: z.string().min(1, "Address is required"),
  identity: z.string().min(1, "Identity is required"),
  phone: z.string().regex(/^\+?[1-9]\d{1,14}$/, "Invalid phone number"),
  description: z.string(),
  problems: z.array(z.string()),
  medicines: z.array(z.string()),
  conditions: z.array(z.string()),
  diagnosis: z.string(),
  next_session: z
    .string()
    .regex(/^\d{4}-\d{2}-\d{2}$/, "Invalid date format (YYYY-MM-DD)")
    .default(new Date().toISOString().split("T")[0])
    .optional(),
});

export type PatientFormData = z.infer<typeof patientSchema>;

interface EMRFormProps {
  initialData?: Partial<PatientFormData>;
}

async function onSubmit(data: PatientFormData) {
  console.log(data);
  const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;
  const finalData: unknown = data;
  finalData["problems"] = JSON.stringify(finalData["problems"]);
  finalData["conditions"] = JSON.stringify(finalData["conditions"]);
  finalData["medicines"] = JSON.stringify(finalData["medicines"]);
  const res = await fetch(`http://${BACKEND_URL}/patients/${data.patient_id}`, {
    method: "PUT",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json",
    },
  });

  const result = await res.json();
  console.log("Patient Updated:", result);
}

export function EMRForm({ initialData }: EMRFormProps) {
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
    <Card className="w-full max-w-4xl mx-auto p-4">
      <CardHeader>
        <CardTitle>Electronic Medical Record</CardTitle>
        <CardDescription>Patient #{initialData.patient_id}</CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
          <div className="grid grid-cols-2 gap-6">
            <div>
              <Label htmlFor="name">Name</Label>
              <Input id="last_name" {...register("name")} />
              {errors.name && (
                <p className="text-red-500 text-sm">{errors.name.message}</p>
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
                    <RadioGroupItem value="Male" id="male" />
                    <Label htmlFor="male">Male</Label>
                  </div>
                  <div className="flex items-center space-x-2">
                    <RadioGroupItem value="Female" id="female" />
                    <Label htmlFor="female">Female</Label>
                  </div>
                  <div className="flex items-center space-x-2">
                    <RadioGroupItem value="Other" id="other" />
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
          <div className="grid grid-cols-2 gap-6">
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
            <Label htmlFor="description">Description</Label>
            <Textarea id="description" {...register("description")} />
            {errors.description && (
              <p className="text-red-500 text-sm">
                {errors.description.message}
              </p>
            )}
          </div>
          <AutoExpandingInputGroup
            control={control}
            name="problems"
            label="Problems"
            error={errors.problems}
          />
          <AutoExpandingInputGroup
            control={control}
            name="conditions"
            label="Conditions"
            error={errors.conditions}
          />
          <AutoExpandingInputGroup
            control={control}
            name="medicines"
            label="Medicines"
            error={errors.medicines}
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
          <Button type="submit" className="w-full mt-4">
            Submit
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
