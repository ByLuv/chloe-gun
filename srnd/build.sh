rm *.a *.o
gcc -c rand.c -o rand.o
ar -crs librand.a rand.o
rm *.o
echo "Done!"
