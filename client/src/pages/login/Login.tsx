import bgImg from "../../assets/img/login-bg.png";

export default function Login() {
  return (
    <div className="w-[1024px] flex flex-column">
      <div className="w-1/2">
        <img src={bgImg} className="object-cover" />
      </div>
      <div className="w-1/2">abc</div>
    </div>
  );
}
