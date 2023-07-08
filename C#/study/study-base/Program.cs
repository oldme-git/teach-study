using System;
using Oldme2;

namespace Oldme
{
    abstract class AbOldme
    {
        abstract public int area(int a);
        
        public void ok (string str)
        {
            Console.WriteLine(str);
        }
    }

    class MyOldme : AbOldme
    {
        public override int area(int a)
        {
            return a + 1;
        }
        public int Add(int x, int y, int z) {  
            MyOldme2 myOldme = new MyOldme2();
            return myOldme.Add(x, y) + z; 
        }
    }

    class IMain {
        static void Main()
        {
            MyOldme myOldme = new MyOldme();
            int i = myOldme.area(1);
            Console.WriteLine(i);
            myOldme.ok("my oldme");
        }
    }
}
