#include<iostream>
#include<vector>

int sum_of(std::vector<int> values)
{
	int sum(0);

	for(int j(0);j < values.size() ; j++)
	{
		//std::cout << j << " " << values[j] << " " << sum << std::endl;
		sum+=values[j];
	}

	return sum;
}


int main(void)
{
	std::vector<int> multiples_of_3;
	std::vector<int> multiples_of_5;
	std::vector<int> multiples_of_3_and_5;

	int limit=1000;

	for(int i(0);i<limit;i++)
	{
		if(i%3==0)
		{
			multiples_of_3.push_back(i);
			multiples_of_3_and_5.push_back(i);
		}
		if(i%5==0)
		{
			multiples_of_5.push_back(i);
			multiples_of_3_and_5.push_back(i);
		}
	}

	std::cout << "Total sum is " << sum_of(multiples_of_3_and_5) << std::endl;

	return 0;
}
