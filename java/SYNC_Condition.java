/* Run this to learn how Condition works */
import java.util.concurrent.locks.*;

public class SYNC_Condition {
    private static int v = 0;
    private static ReentrantLock lock = null;
    private static Condition condi = null;
    public static void main(String args[]) {
        lock = new ReentrantLock();
        condi = lock.newCondition();
        System.out.println("Hello");

        Thread t1 = new Thread(new Runnable() {
            public void run() {
                lock.lock();
                try {
                    System.out.println(String.format("[%d] Entering... v=%d", System.currentTimeMillis(), v));
                    while(v < 100) {
                        condi.await();
                        
                        System.out.println(String.format("[%d] Wakeup... v=%d", System.currentTimeMillis(), v));
                    }
                    
                    System.out.println(String.format("[%d] Exiting... v=%d", System.currentTimeMillis(), v));
                }
                catch (InterruptedException e) {}
                finally {
                    lock.unlock();
                }
            }
        });
    

        Thread t2 = new Thread(new Runnable() {
            public void run() {
                while(v < 200) {
                    lock.lock();
                    try {
                        System.out.println(String.format("[%d] about to add: v=%d", System.currentTimeMillis(), v));
                        v += 20;
                        condi.signalAll();
                        System.out.println(String.format("[%d] signaled: v=%d", System.currentTimeMillis(), v));
                    }
                    finally {
                        lock.unlock();
                    }
                    try {
                        Thread.sleep(1000);
                    }
                    catch(Exception e) {}
                }
            }
        });
        t1.start();
        t2.start();
        try {
            t1.join();
        }
        catch (InterruptedException e) {}
        
        try {
            t2.join();
        }
        catch (InterruptedException e) {}
       

    }
}
