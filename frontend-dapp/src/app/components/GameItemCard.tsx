type GameItem = {
  id: string;
  name: string;
  image: string;
  description: string;
  owner: string;
};

export default function GameItemCard({ item }: { item: GameItem }) {
  return (
    <div className="bg-zinc-800 rounded-xl shadow-md p-4 w-full sm:w-64">
      <img src={item.image} alt={item.name} className="rounded-md mb-4" />
      <h3 className="text-xl font-semibold">{item.name}</h3>
      <p className="text-sm text-gray-400">{item.description}</p>
      <p className="text-xs mt-2 text-purple-400 truncate">
        Owned by: {item.owner}
      </p>
    </div>
  );
}
