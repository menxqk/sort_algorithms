
void bubble_sort(int arr[], int size)
{
    int i, j;
    int temp;

    for (i = 0; i < size - 1; i++)
    {
        for (j = 0; j < size - i - 1; j++)
        {
            if (arr[j] > arr[j + 1])
            {
                temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
}

void selection_sort(int arr[], int size)
{
    int i, j, min_idx;
    int temp;

    for (i = 0; i < size - 1; i++)
    {
        min_idx = i;

        for (j = i + 1; j < size; j++)
        {
            if (arr[j] < arr[min_idx])
            {
                min_idx = j;
            }
        }

        if (min_idx != i)
        {
            temp = arr[i];
            arr[i] = arr[min_idx];
            arr[min_idx] = temp;
        }
    }
}

void insertion_sort(int arr[], int size)
{
    int i, key, j;

    for (i = 1; i < size; i++)
    {
        key = arr[i];
        j = i - 1;

        while (j >= 0 && arr[j] > key)
        {
            arr[j + 1] = arr[j];
            j = j - 1;
        }

        arr[j + 1] = key;
    }
}

static void merge(int arr[], int beg, int mid, int end)
{
    int i, j, k;
    int n1 = mid - beg + 1;
    int n2 = end - mid;

    int left_array[n1];
    int right_array[n2];

    for (i = 0; i < n1; i++)
    {
        left_array[i] = arr[beg + i];
    }

    for (j = 0; j < n2; j++)
    {
        right_array[j] = arr[mid + 1 + j];
    }

    i = 0;
    j = 0;
    k = beg;

    while (i < n1 && j < n2)
    {
        if (left_array[i] <= right_array[j])
        {
            arr[k] = left_array[i];
            i++;
        }
        else
        {
            arr[k] = right_array[j];
            j++;
        }
        k++;
    }

    while (i < n1)
    {
        arr[k] = left_array[i];
        i++;
        k++;
    }

    while (j < n2)
    {
        arr[k] = right_array[j];
        j++;
        k++;
    }
}

void merge_sort(int arr[], int beg, int end)
{
    if (beg < end)
    {
        int mid = (beg + end) / 2;
        merge_sort(arr, beg, mid);
        merge_sort(arr, mid + 1, end);
        merge(arr, beg, mid, end);
    }
}

static void heapify(int arr[], int n, int i)
{
    int largest = i;
    int left = 2 * i + 1;
    int right = 2 * i + 2;
    int temp;

    if (left < n && arr[left] > arr[largest])
    {
        largest = left;
    }

    if (right < n && arr[right] > arr[largest])
    {
        largest = right;
    }

    if (largest != i)
    {
        temp = arr[i];
        arr[i] = arr[largest];
        arr[largest] = temp;

        heapify(arr, n, largest);
    }
}

void heap_sort(int arr[], int size)
{
    int i;
    int temp;

    for (i = size / 2 - 1; i >= 0; i--)
    {
        heapify(arr, size, i);
    }

    for (i = size - 1; i >= 0; i--)
    {
        temp = arr[0];
        arr[0] = arr[i];
        arr[i] = temp;

        heapify(arr, i, 0);
    }
}

static int partition(int arr[], int beg, int end)
{
    int pivot = arr[end];
    int i = beg - 1;
    int j;
    int temp;

    for (j = beg; j <= end - 1; j++)
    {
        if (arr[j] < pivot)
        {
            i++;
            temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
        }
    }

    temp = arr[i + 1];
    arr[i + 1] = arr[end];
    arr[end] = temp;
    return i + 1;
}

void quick_sort(int arr[], int beg, int end)
{
    if (beg < end)
    {
        int p = partition(arr, beg, end);
        quick_sort(arr, beg, p - 1);
        quick_sort(arr, p + 1, end);
    }
}