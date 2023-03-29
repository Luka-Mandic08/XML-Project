import styles from './add-flight-container.module.css';
import { useForm } from "react-hook-form";

type FormValues = {
  startdate       :string,        
	arrivaldate     :string,     
	destination     :string,     
	start           :string,   
	price           :number,
	totaltickets    :number,
}

/* eslint-disable-next-line */
export interface AddFlightContainerProps {}

export function AddFlightContainer(props: AddFlightContainerProps){
  const { register, handleSubmit, formState: {errors} } = useForm(
    {
      defaultValues: {
        startdate: Date.now().toString(),
        arrivaldate: Date.now.toString(),
        destination: "",
        start: "",
        price: 1,
        totaltickets: 1

      }
    }
  );
  const onSubmit = (data: FormValues) => console.log(data);


  return (
    <div className={styles.box}>
    <form onSubmit={handleSubmit(onSubmit)}>
      <input
        type="date" 
        {...register("startdate", {required: 'This field is required.'})}

      />
      <p>{errors.startdate?.message}</p>
      <input
        type="date" 
        {...register("arrivaldate", {required: 'This field is required.'})}

      />
      <p>{errors.arrivaldate?.message}</p>
      <div className={styles.inputContainer}>
        <input
          type="text" 
          {...register("destination", {required: 'This field is required.'})}
          placeholder="Destination"
        />
        <p>{errors.destination?.message}</p>
      </div>
      <input
        type="text" 
        {...register("start", {required: 'This field is required.'})}
        placeholder="Starting location"
      />
      <p>{errors.start?.message}</p>
      <input
        type="number" 
        {...register("price", {required: 'This field is required.', min:{
          value:1,
          message: "Minimal price is 1 ."
        } })}
        placeholder="Price"
      />
      <p>{errors.price?.message}</p>
      <input
        type="number" 
        {...register("totaltickets", {required: 'This field is required.', min:{
          value:1,
          message: "Minimal number of passangers is 1."
        }})}
        placeholder="Number of passangers"
      />
      <p>{errors.totaltickets?.message}</p>
      <input
        type="submit"
      />
    </form>
    </div>
  );
}

export default AddFlightContainer;
