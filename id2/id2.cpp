#include<iostream>
#include<vector>


int fibonacci_n(int Nth)
{
	int fibonacci_number(1);
	int fibonacci_number_b(1);
	int fibonacci_seed(1);

	for(int i=1;i<Nth;i++)
	{
		fibonacci_number_b = fibonacci_number + fibonacci_seed;
		fibonacci_seed = fibonacci_number;
		fibonacci_number = fibonacci_number_b;
	}
	return fibonacci_number;
}


int main(void)
{
	//int limit(10);
	int limit(4000000);
	int sum_tot(0);
	int iBuffer(0);
	int counter(0);

	for(int i(1); i<=limit; i++)
	{
	    iBuffer = fibonacci_n(i);
		
		//std::cout << i << "\t" << iBuffer << std::endl;

		if(iBuffer%2 == 0 && iBuffer < limit )
		{
			sum_tot+=fibonacci_n(i);
			counter+=1;
			//std::cout << fibonacci_n(i) << " " << sum_tot << std::endl;
		}
		else if( iBuffer >= limit)
		{
			break;
		}
	}

	std::cout << "Total sum of first " << counter << " even Fibonacci numbers is " << sum_tot << std::endl;

	return 0;
}
