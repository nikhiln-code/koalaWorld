export default function Button({
  children,
  onClick,
  disabled,
  className,
}: {
  children: ReactNode;
  onClick?: () => void;
  disabled?: boolean;
  className?: string;
}) {
  return (
    <button
      onClick={onClick}
      disabled={disabled}
      className={`px-6 py-3 bg-gradient-to-r from-green-500 via-lime-600 to-green-700 text-white font-bold rounded-xl shadow-lg hover:scale-105 transition duration-300 ${className}`}
    >
      {children}
    </button>
  );
}
