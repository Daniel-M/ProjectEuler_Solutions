#include<iostream>
#include<vector>

int main(void)
{
	int limit(1000);
	int sum(0);

	for(int i(0);i<limit;i++)
	{
		if((i%3==0) || (i%5==0))
		{
			sum+=i;
		}
	}

	std::cout << "Total sum of multiples of 3 or 5 is" << sum << std::endl;

	return 0;
}
