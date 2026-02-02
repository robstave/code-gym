type City = {
	name: string;
	state: string;
	population: number;
};

const cities: City[] = [
	{ name: "Springfield", state: "IL", population: 116250 },
	{ name: "Springfield", state: "MA", population: 153606 },
	{ name: "Austin", state: "TX", population: 964254 },
	{ name: "Springfield", state: "MO", population: 169176 },
	{ name: "Columbus", state: "OH", population: 905748 },
	{ name: "Austin", state: "MN", population: 25000 },
	{ name: "Riverside", state: "CA", population: 331360 },
	{ name: "Columbus", state: "GA", population: 195769 },
];

// Sort by city (case-insensitive). If city names collide, secondary sort by population (descending).
function sortCities(list: City[]): City[] {

    list.sort((a,b) => {

        if (a.name > b.name) {
            return 1
        } else if ( a.name < b.name) {
            return -1
        } else {
            return (b.population - a.population)
        }

    })
 
     return list
}

const sorted = sortCities(cities);
console.log(sorted);