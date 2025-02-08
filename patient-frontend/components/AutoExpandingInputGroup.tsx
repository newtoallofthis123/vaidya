import { useFieldArray, type Control } from "react-hook-form";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { PlusCircle, X } from "lucide-react";

interface AutoExpandingInputGroupProps {
  control: Control<any>;
  name: string;
  label: string;
  error?: any;
}

export function AutoExpandingInputGroup({
  control,
  name,
  label,
  error,
}: AutoExpandingInputGroupProps) {
  const { fields, append, remove } = useFieldArray({
    control,
    name,
  });

  return (
    <div>
      <Label>{label}</Label>
      {fields.map((field, index) => (
        <div key={field.id} className="flex items-center space-x-2 mt-2">
          <Input
            {...control.register(`${name}.${index}`)}
            defaultValue={field.value}
          />
          <Button
            type="button"
            variant="ghost"
            size="icon"
            onClick={() => remove(index)}
          >
            <X className="h-4 w-4" />
          </Button>
        </div>
      ))}
      <Button
        type="button"
        variant="outline"
        size="sm"
        className="mt-2"
        onClick={() => append("")}
      >
        <PlusCircle className="h-4 w-4 mr-2" />
        Add {label}
      </Button>
      {error && <p className="text-red-500 text-sm mt-1">{error.message}</p>}
    </div>
  );
}
